from flask import render_template, url_for, flash, redirect, request, abort, Blueprint
from app_flask.models import User
from app_flask import db, bcrypt
from flask_login import current_user, login_required

from app_flask.users.utils import (
    get_all_users)
from app_flask.users.forms import (newUserForm)

users = Blueprint('users', __name__)


@users.route('/')
@login_required
def users_home():
    return redirect(url_for('users.users_management'))


@users.route('/users')
@login_required
def users_management():
    users_populated = get_all_users()
    return render_template('users.html', title='Users Management',
                           users=users_populated,
                           current_user=current_user,
                           active_page='users')


@users.route("/users/new", methods=['GET', 'POST'])
@login_required
def user_new():
    user_new_form = newUserForm()
    if request.method == 'POST' and user_new_form.validate_on_submit():
        hashed_password = bcrypt.generate_password_hash(
            user_new_form.password.data).decode('utf-8')
        user = User(name=user_new_form.name.data,
                    email=user_new_form.email.data,
                    password=hashed_password)
        db.session.add(user)
        db.session.commit()

        flash(f"The user {user.name} has been created!.", 'success')
        return redirect(url_for('users.users_management'))

    return render_template('user_new.html',
                           title='User New',
                           form=user_new_form,
                           save_name='Create User',
                           save_icon='fa-plus',
                           current_user=current_user)
