"""
Setting easy to read/maintain variables:

-----------
Web Variant
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
-----------
Only project with web:
{% if cookiecutter.web_variant == 'y' %}
    {% if cookiecutter.css_addon == 'Bootstrap' %}
        {{ cookiecutter.update({ "bootstrap_addon": "y" }) }}
        {{ cookiecutter.update({ "tailwinds_addon": "n" }) }}
    {% else %}
        {% if cookiecutter.css_addon == 'Tailwinds' %}
            {{ cookiecutter.update({ "bootstrap_addon": "n" }) }}
            {{ cookiecutter.update({ "tailwinds_addon": "y" }) }}
        {% else %}
            {{ cookiecutter.update({ "bootstrap_addon": "n" }) }}
            {{ cookiecutter.update({ "tailwinds_addon": "n" }) }}
        {% endif %}
    {% endif %}
Ignored for API only projects
{% else %}
    {{ cookiecutter.update({ "bootstrap_addon": "n" }) }}
    {{ cookiecutter.update({ "tailwinds_addon": "n" }) }}
{% endif %}
"""
