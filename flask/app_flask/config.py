import os

class Config:
    def __init__(self):
        self.SECRET_KEY = os.environ.get('SECRET_KEY')
        
        self.SQLALCHEMY_ECHO = True  # This enables logging
        self.SQLALCHEMY_TRACK_MODIFICATIONS = True  # This enables logging

        parent_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '..', '..'))
        database_dir = os.path.join(parent_dir, 'instance')
        database_file = os.path.join(database_dir, os.environ.get('SQLALCHEMY_DATABASE_URI'))
        
        self.SQLALCHEMY_DATABASE_URI = 'sqlite:///' + database_file


