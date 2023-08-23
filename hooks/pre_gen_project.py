'''
Setting easy to read/maintain cookiecutter variables:

-----------
Web/API Variant
cookiecutter._web_variant
cookiecutter._api_variant
-----------
{% if cookiecutter.variant in ['Web only', 'Both'] %}
    {{ cookiecutter.update({ '_web_variant': 'yes' }) }}
{% endif %}

{% if cookiecutter.variant in ['API only', 'Both'] %}
    {{ cookiecutter.update({ '_api_variant': 'yes' }) }}
{% endif %}

-----------
CSS Addons
cookiecutter._bootstrap_addon
cookiecutter._tailwind_addon
-----------
Only project with web:
{% if cookiecutter._web_variant == 'yes' %}
    {% if cookiecutter.css_addon == 'Bootstrap' %}
        {{ cookiecutter.update({ '_bootstrap_addon': 'yes' }) }}
    {% elif cookiecutter.css_addon == 'Tailwind' %}
        {{ cookiecutter.update({ '_tailwind_addon': 'yes' }) }}
    {% endif %}
{% endif %}
'''
