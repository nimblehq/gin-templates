"""
Setting easy to read/maintain cookiecutter variables:

-----------
Web Variant
cookiecutter.web_addon
-----------
{% if cookiecutter.variant in ["Web only", "Both"] %}
    {{ cookiecutter.update({ "web_variant": True }) }}
{% else %}
    {{ cookiecutter.update({ "web_variant": False }) }}
{% endif %}

API Variant
{% if cookiecutter.variant in ["API only", "Both"] %}
    {{ cookiecutter.update({ "api_variant": True }) }}
{% else %}
    {{ cookiecutter.update({ "api_variant": False }) }}
{% endif %}

-----------
CSS Addon
cookiecutter.css_addon
-----------
Ignored unless a condition matches
{{ cookiecutter.update({ "bootstrap_addon": False }) }}
{{ cookiecutter.update({ "tailwinds_addon": False }) }}

Only project with web:
{% if cookiecutter.web_variant %}
    {% if cookiecutter.css_addon == "Bootstrap" %}
        {{ cookiecutter.update({ "bootstrap_addon": True }) }}
    {% elif cookiecutter.css_addon == "Tailwinds" %}
        {{ cookiecutter.update({ "tailwinds_addon": True }) }}
    {% endif %}
{% endif %}
"""
