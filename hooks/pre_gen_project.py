'''
Setting easy to read/maintain cookiecutter variables:

-----------
Web/API Variant
cookiecutter._web_variant
cookiecutter._api_variant
-----------
{% if cookiecutter.variant in ['Web only', 'Both'] %}
    {{ cookiecutter.update({ '_web_variant': 'y' }) }}
{% endif %}

API Variant
{% if cookiecutter.variant in ['API only', 'Both'] %}
    {{ cookiecutter.update({ '_api_variant': 'y' }) }}
{% endif %}

-----------
CSS Addons
cookiecutter._bootstrap_addon
cookiecutter._tailwinds_addon
-----------

Only project with web:
{% if cookiecutter._web_variant == 'y' %}
    {% if cookiecutter.css_addon == 'Bootstrap' %}
        {{ cookiecutter.update({ '_bootstrap_addon': 'y' }) }}
    {% elif cookiecutter.css_addon == 'Tailwinds' %}
        {{ cookiecutter.update({ '_tailwinds_addon': 'y' }) }}
    {% endif %}
{% endif %}
'''
