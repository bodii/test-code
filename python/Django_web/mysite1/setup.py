import os
from setuptools import find_packages, setup

with open(os.path.join(os.path.dirname(__file__), 'README.rst')) as readme:
    README = readme.read()

    # allow setup.py to be run from any path
    os.chdir(os.path.normpath(os.path.join(os.path.abspath(__file__), os.pardir)))

    setup(
        name = 'django-polls',
        version = '0.1',
        packages = find_packages(),
        include_package_data = True,
        license = 'BSD License',
        description = 'A simple Django app to conduct Web-based polls.',
        long_description = README,
        url = 'https://www.example.com/',
        author = 'Your Name',
        author_email = 'yourname@example.com',
        classifiers = [
            'Environment :: Web Environment',
            'Framework :: Django',
            'Framework :: Django :: X.Y',
            'Intended Audience :: Developers',
            'License :: OSI Approved :: BSD License',
            'Operating System :: OS independent',
            'Programming Language :: Python',
            'Programming Language :: Python :: 3.6',
            'Topic :: Internet :: WWW/HTTP',
            'Topic :: Interent :: WWW/HTTP :: Dynamic Content',
        ],
    )