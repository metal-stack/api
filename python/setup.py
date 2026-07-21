from setuptools import setup, find_packages
import os

NAME = "metal-stack-api"

REQUIRES = [
    # this version needs to match the version specified in buf.gen.yaml
    "connectrpc==0.10.1",
    "protobuf>=7.0",
    "protovalidate>=1.2.0",
    "pyqwest<0.7.0",
]

setup(
    name=NAME,
    version=os.environ["VERSION"],
    description="Python API client for metal-stack api",
    long_description="Python API client for metal-stack api that implements the v2 api and deprecates metal_python.",
    author="metal-stack authors",
    url="https://github.com/metal-stack/api",
    keywords=["metal-stack", "metal-apiserver"],
    install_requires=REQUIRES,
    license="MIT",
    packages=find_packages(),
    classifiers=[
        'Intended Audience :: Developers',
        'Natural Language :: English',
        'Operating System :: POSIX',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.13',
        'Programming Language :: Python :: 3.14',
    ],
    include_package_data=True,
)
