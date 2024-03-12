from flask_wtf import FlaskForm
from wtforms import StringField, PasswordField, SubmitField
from wtforms.validators import DataRequired, Length, Email, ValidationError

from app_flask.models import User

class newUserForm(FlaskForm):
    name = StringField('Nombre',
                           validators=[DataRequired(), Length(min=2, max=20)])
    email = StringField('Correo Electrónico',
                        validators=[DataRequired(), Email()])
    password = PasswordField('Contraseña',
                           validators=[DataRequired(), Length(min=5, max=20)])
    submit = SubmitField('Crear Usuario')

    def validate_name(self, name):
        user = User.query.filter_by(name=name.data).first()
        if user:
            raise ValidationError(
                'Nombre no disponible, escoge uno diferente.')

    def validate_email(self, email):
        user = User.query.filter_by(email=email.data).first()
        if user:
            raise ValidationError(
                'Email no disponible, escoge uno diferente.')
