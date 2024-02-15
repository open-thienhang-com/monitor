from setuptools import setup, find_packages

setup(
    name='dev.thienhang.com',                # Replace with your project's name
    version='0.1',                   # Replace with your project's version
    description='My Python project', # Replace with a project description
    author='Your Name',              # Replace with your name
    author_email='you@example.com',  # Replace with your email
    py_modules=['main'],
    packages=find_packages(),        # Automatically find and include all packages
    install_requires=[
        'requests',                  # List your project's dependencies here
        # Other dependencies...
    ],
)
