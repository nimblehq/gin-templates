"""
Setting easy to read/maintain cookiecutter variables:

-----------
Web Variant
cookiecutter.web_addon
-----------
{% if cookiecutter.variant in ['Web only', 'Both'] %}
    {{ cookiecutter.update({ "web_variant": "y" }) }}
{% else %}
    {{ cookiecutter.update({ "web_variant": "n" }) }}
{% endif %}

API Variant
{% if cookiecutter.variant in ['API only', 'Both'] %}
    {{ cookiecutter.update({ "api_variant": "y" }) }}
{% else %}
    {{ cookiecutter.update({ "api_variant": "n" }) }}
{% endif %}

-----------
CSS Addon
cookiecutter.css_addon
-----------
Ignored unless a condition matches
{{ cookiecutter.update({ "bootstrap_addon": "n" }) }}
{{ cookiecutter.update({ "tailwinds_addon": "n" }) }}

Only project with web:
{% if cookiecutter.web_variant == 'y' %}
    {% if cookiecutter.css_addon == 'Bootstrap' %}
        {{ cookiecutter.update({ "bootstrap_addon": "y" }) }}
    {% elif cookiecutter.css_addon == 'Tailwinds' %}
        {{ cookiecutter.update({ "tailwinds_addon": "y" }) }}
    {% endif %}
{% endif %}
"""
