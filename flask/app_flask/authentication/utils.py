from app_flask import bcrypt

def check_password(user, password_eval):
    return bcrypt.check_password_hash(user.password, password_eval)