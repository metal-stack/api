from setuptools import setup, find_packages
import os

NAME = "metal-stack-api"

REQUIRES = [
    "connecpy>=1.5.2",
]

setup(
    name=NAME,
    version=os.environ.get("VERSION", "v0.0.1"),
    description="Python API client for metal-stack api (v2)",
    author="metal-stack authors",
    url="https://github.com/metal-stack/api",
    keywords=["metal-stack", "metal-apiserver"],
    install_requires=REQUIRES,
    license="MIT",
    packages=find_packages(),
    classifiers=[
        "License :: OSI Approved :: MIT License",
        'Intended Audience :: Developers',
        "License :: OSI Approved :: MIT License",
        'Natural Language :: English',
        'Operating System :: POSIX',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.11',
        'Programming Language :: Python :: 3.12',
        'Programming Language :: Python :: 3.13',
    ],
    include_package_data=True,
)
