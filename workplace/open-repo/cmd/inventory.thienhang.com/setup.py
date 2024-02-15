from setuptools import setup, find_packages

setup(
    name='inventory.thienhang.com',                # Replace with your project's name
    version='0.1',                   # Replace with your project's version
    description='Inventory', # Replace with a project description
    author='inventory.thienhang.com',              # Replace with your name
    author_email='me@thienhang.com',  # Replace with your email
    py_modules=['main'],
    packages=find_packages(),        # Automatically find and include all packages
    install_requires=[
        'requests',                  # List your project's dependencies here
        # Other dependencies...
    ],
)
