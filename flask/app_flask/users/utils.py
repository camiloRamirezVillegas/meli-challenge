from app_flask.models import User
from app_flask import db


# @login_manager.user_loader
# def load_user(user_id):
#     current_user_load = User.query.get(int(user_id))
#     # print("current_user_load: ", current_user_load)

#     return current_user_load


def get_all_users():
    #     users_list = db.session.query(User)\
    #     .join(Client, (Client.id == User.client_id)).all()
    # return users_list

    users_list = db.session.query(User)\
        .all()
    return users_list