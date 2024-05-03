from .base import *

# from mongoengine import connect

# Database
# https://docs.djangoproject.com/en/4.2/ref/settings/#databases

MONGODB_DATABASES = {
    "default": {
        "name": "happy-wallet",
        "host": "mongodb+srv://shccgxqp:Qazws12345@googletw.zcugz3d.mongodb.net/",
        "tz_aware": True,  # if you using timezones in django (USE_TZ = True)
    },
}


ALLOWED_HOSTS = ["*"]
