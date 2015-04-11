#!/usr/bin/env python
# -*- coding: utf-8 -*-


import json
import os


import falcon


from . import config


class Item(object):

    def on_get(self, req, resp, service, configure):
        if not os.path.join(config.STORE_PATH, service, configure + ".json"):
            raise falcon.HTTPNotFound
        content = open(os.path.join(config.STORE_PATH, service, configure + ".json"), 'rb').read()
        resp.body = json.dumps({
            'status': "success",
            'data': json.loads(content),
        })

    def on_post(self, req, resp, service, configure):
        try:
            conf = open(os.path.join(config.STORE_PATH, service, configure + ".json"), 'w')
        except IOError:
            raise falcon.HTTPNotFound
        conf.write(req.stream.read())
        resp.body = json.dumps({
            'status': "success",
        })
