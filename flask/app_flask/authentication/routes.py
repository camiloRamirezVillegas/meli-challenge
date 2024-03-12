from flask import render_template, url_for, flash, redirect, request, Blueprint
from flask_login import login_user, current_user, logout_user

from app_flask.authentication.forms import (LoginForm)
from app_flask.authentication.utils import (check_password)
from app_flask.users.utils import (get_user_by_email)

authentication = Blueprint('authentication', __name__)

@authentication.route("/login", methods=['GET', 'POST'])
def login():
    form = LoginForm()

    if current_user.is_authenticated:
        return redirect(url_for('users.users_management'))
        
    if form.validate_on_submit():
        user = get_user_by_email(email=form.email.data)

        if user and check_password(user, form.password.data):
            login_user(user, remember=form.remember.data)
            next_page = request.args.get('next')
            return redirect(next_page) if next_page else redirect(url_for('users.users_management'))
        else:
            flash('Inicio de sesión fallido. Revisa el email y/o contraseña.', 'danger')

    return render_template('login.html', title='Login', form=form)

@authentication.route("/logout")
def logout():
    logout_user()
    return redirect(url_for('authentication.login'))
