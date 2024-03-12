from app_flask import db
from flask_login import UserMixin
from app_flask import login_manager


# TODO, mover esto?
@login_manager.user_loader
def load_user(user_id):
    current_user_load = User.query.get(int(user_id))
    return current_user_load

#TODO, poner edad, genero, direccion, profesión, fecha nacimiento ?, telefono
class User(db.Model, UserMixin):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(255), unique=True, nullable=False)
    email = db.Column(db.String(255), unique=True, nullable=False)
    password = db.Column(db.String(255), nullable=False)

    def __repr__(self):
        return f"User('{self.name}', '{self.email}')"

    @property
    def serialize(self):
        """Return object data in easily serializable format"""
        return {
            'id': self.id,
            'name': self.name,
            'email': self.email,
        }

