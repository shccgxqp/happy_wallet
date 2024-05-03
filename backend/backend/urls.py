from django.contrib import admin
from django.urls import path, include
from rest_framework.routers import DefaultRouter

# from backend.api import views

# router = DefaultRouter()
# router.register(r"users", views.UserViewSet)

urlpatterns = [
    path("admin/", admin.site.urls),
    # path("api/v1/", include(router.urls)),
]
