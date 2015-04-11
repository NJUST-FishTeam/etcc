#!/usr/bin/env python
# -*- coding: utf-8 -*-


import falcon


from . import service

api = application = falcon.API()

api.add_route('/', service.Collection())
api.add_route('/{service}', service.Item())
