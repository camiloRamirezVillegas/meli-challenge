from app_flask.models import User
from app_flask import db, bcrypt

def get_all_users():
    users_list = db.session.query(User).all()
    return users_list

def get_user_by_email(email):
    user = User.query.filter_by(email=email).first()
    return user

def create_user(user_form):
    hashed_password = bcrypt.generate_password_hash(user_form.password.data).decode('utf-8')
    user = User(name=user_form.name.data,
                email=user_form.email.data,
                password=hashed_password)
    db.session.add(user)
    db.session.commit()