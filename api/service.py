#!/usr/bin/env python
# -*- coding: utf-8 -*-


import json
import os


import falcon


from . import config


class Collection(object):

    def on_get(self, req, resp):
        files = os.listdir(config.STORE_PATH)
        services = [f for f in files
                    if os.path.isdir(os.path.join(config.STORE_PATH, f))]
        resp.body = json.dumps({
            'status': 'success',
            'data': services,
        })


class Item(object):

    def on_get(self, req, resp, service):
        service_path = os.path.join(config.STORE_PATH, service)
        if not os.path.exists(service_path):
            raise falcon.HTTPNotFound
        files = os.listdir(service_path)
        configs = [f.split('.')[0] for f in files
                   if os.path.isfile(os.path.join(service_path, f))]
        data = {}
        for conf in configs:
            content = open(os.path.join(service_path, conf + '.json')).read()
            data[conf] = json.loads(content)

        resp.body = json.dumps({
            'status': 'success',
            'data': data
        })
