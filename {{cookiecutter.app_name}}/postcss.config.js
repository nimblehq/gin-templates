module.exports = {
  plugins: [
    require("autoprefixer"){% if cookiecutter._tailwind_addon == "yes" %},
    require("tailwindcss"){% endif %}
  ],
};
