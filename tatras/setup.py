from setuptools import setup

setup(
    name='tatras',
	version='1.1',
	description='ArgoCD Extension',
	author='Jack Harmon',
	author_email='john.harmon96@gmail.com',
	packages=['tatras'],
    url='https://github.com/tatrasproject/tatras',
	install_requires=[
        'wheel', 'requests', 'python-dotenv', 
        'fastapi', 'uvicorn[standard]'
        ]
)