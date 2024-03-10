from flask import render_template, url_for, flash, redirect, request, abort, Blueprint
from flask_login import login_user, current_user, logout_user, login_required

from app_flask.authentication.forms import (LoginForm)
from app_flask.models import User
from app_flask import db, bcrypt

authentication = Blueprint('authentication', __name__)

# TODO, mover las logicas complejas a donde corresponda, por ejemplo la consulta del usuario deberia salir del modulo de users
@authentication.route("/login", methods=['GET', 'POST'])
def login():
    if current_user.is_authenticated:
        return redirect(url_for('users.users_management'))
    form = LoginForm()
    if form.validate_on_submit():
        user = User.query.filter_by(email=form.email.data).first()
        if user and bcrypt.check_password_hash(user.password, form.password.data):
            login_user(user, remember=form.remember.data)
            next_page = request.args.get('next')
            return redirect(next_page) if next_page else redirect(url_for('users.users_management'))
        else:
            flash('Login Unsuccessful. Please check email and password', 'danger')
    return render_template('login.html', title='Login', form=form)

@authentication.route("/logout")
def logout():
    logout_user()
    return redirect(url_for('users.users_management'))
