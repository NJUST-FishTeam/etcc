#!/usr/bin/env python
# -*- coding: utf-8 -*-


import json
import os


import falcon


from . import config


class Collection(object):
    pass


class Item(object):

    def on_get(self, req, resp, service, configure):
        if not os.path.join(config.STORE_PATH, service, configure + ".json"):
            raise falcon.HTTPNotFound
        content = open(os.path.join(config.STORE_PATH, service, configure + ".json"), 'rb').read()
        resp.body = json.dumps({
            'status': "success",
            'data': json.loads(content),
        })

    def on_put(self, req, resp, service, configure):
        conf_path = os.path.join(config.STORE_PATH, service, configure + '.json')
        if not os.path.exists(conf_path):
            raise falcon.HTTPNotFound
        conf = open(conf_path, 'w')
        conf.write(req.stream.read())
        resp.body = json.dumps({
            'status': "success",
        })
