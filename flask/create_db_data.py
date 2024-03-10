import random
from app_flask.models import User
from app_flask import db, create_app, bcrypt

from app_flask.models import User

app = create_app()
app.app_context().push()
db.drop_all()
db.create_all()

user_pass = bcrypt.generate_password_hash('54321').decode('utf-8')

first_user = User(name='Camilo Ramirez',
                    email='camiramirezv@gmail.com',
                    password=user_pass)
db.session.add(first_user)
db.session.commit()

