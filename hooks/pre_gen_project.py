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

API Variant
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

-----------
Repository Addons
cookiecutter._github_addon
cookiecutter._gitlab_addon
cookiecutter._gitignore_addon
-----------
{% if cookiecutter.repository_addon == 'GitHub' %}
    {{ cookiecutter.update({ '_github_addon': 'yes' }) }}
{% elif cookiecutter.repository_addon == 'GitLab' %}
    {{ cookiecutter.update({ '_gitlab_addon': 'yes' }) }}
{% elif cookiecutter.repository_addon == 'None' %}
    {{ cookiecutter.update({ '_gitignore_addon': 'no' }) }}
{% endif %}
'''
