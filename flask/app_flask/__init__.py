from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from dotenv import load_dotenv
from flask_bcrypt import Bcrypt
from flask_login import LoginManager

load_dotenv()  # take environment variables from .env.
from app_flask.config import Config  # NOQA for avoiding re order of import, autopep8

db = SQLAlchemy()
config_class = Config()
bcrypt = Bcrypt()
login_manager = LoginManager()

from app_flask.users.routes import users
from app_flask.authentication.routes import authentication

login_manager.login_view = 'authentication.login'
login_manager.login_message_category = 'info'

config_class = Config()

def create_app(config_class=config_class):
    app = Flask(__name__, instance_relative_config=True)
    app.config.from_object(config_class)

    db.init_app(app)
    bcrypt.init_app(app)
    login_manager.init_app(app)

    app.register_blueprint(users)
    app.register_blueprint(authentication)

    return app

