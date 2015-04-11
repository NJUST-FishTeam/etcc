#!/usr/bin/env python
# -*- coding: utf-8 -*-


import json
import os


import falcon


from . import config


class Collection(object):

    def on_get(self, req, resp, service):
        service_path = os.path.join(config.STORE_PATH, service)
        if not os.path.exists(service_path):
            raise falcon.HTTPNotFound
        files = os.listdir(service_path)
        configs = [f.split('.')[0] for f in files
                   if os.path.isfile(os.path.join(service_path, f))]

        data = []
        for conf in configs:
            conf_path = os.path.join(service_path, conf + '.json')
            data.append({
                'name': conf,
                'data': json.loads(open(conf_path).read())
            })

        resp.body = json.dumps({
            'status': 'success',
            'data': data,
            'count': len(configs)
        })

    def on_post(self, req, resp, service):
        service_path = os.path.join(config.STORE_PATH, service)
        if not os.path.exists(service_path):
            raise falcon.HTTPNotFound
        body = json.loads(req.stream.read())
        if body['action'] != 'create_configure':
            resp.status = falcon.HTTP_400
            resp.body = json.dumps({
                'status': "failed",
                'reason': "action is not allowed"
            })
            return
        conf_path = os.path.join(config.STORE_PATH, service, body['configure_name'] + '.json')
        if os.path.exists(conf_path):
            resp.status = falcon.HTTP_400
            resp.body = json.dumps({
                'status': 'failed',
                'reason': 'new configure exists'
            })
            return
        conf = open(conf_path, 'w')
        conf.write(json.dumps(body['configure']))
        resp.body = json.dumps({
            'status': "success",
        })


class Item(object):

    def on_get(self, req, resp, service, configure):
        conf_path = os.path.join(config.STORE_PATH, service, configure + ".json")
        if not os.path.exists(conf_path):
            raise falcon.HTTPNotFound
        content = open(conf_path, 'rb').read()
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

    def on_delete(self, req, resp, service, configure):
        conf_path = os.path.join(config.STORE_PATH, service, configure + '.json')
        if not os.path.exists(conf_path):
            raise falcon.HTTPNotFound
        os.remove(conf_path)
        resp.body = json.dumps({
            'status': "success",
        })
