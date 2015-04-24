#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Author: Comzyh
# @Date:   2015-04-24 01:29:56
# @Last Modified by:   Comzyh
# @Last Modified time: 2015-04-24 03:09:16

import falcon
import os
import mimetypes

class ClassName(object):
    """docstring for ClassName"""


class StaticFiles(object):
    def __init__(self):
        super(StaticFiles, self).__init__()
        etcc_path = os.path.dirname(os.path.dirname(os.path.realpath(__file__)))
        self.static_path = os.path.join(etcc_path,'frontend/static/')

    def on_get(self, req, resp):
        filepath = os.path.join(self.static_path, req.path[8:])
        # resp.body = self.static_path + '\n' + filepath
        if (os.path.isfile(filepath)):
            resp.status = falcon.HTTP_200
            resp.stream = open(filepath, 'rb')
            resp.stream_len = os.path.getsize(filepath)
            resp.content_type, encoding = mimetypes.guess_type(req.url)
            resp.append_header('req.url',mimetypes.guess_type(req.url))
            resp.cache_control=['max-age=30']
        else:
            resp.status = falcon.HTTP_404

class DefaultPage(object):
    def __init__(self):
        super(DefaultPage, self).__init__()
        etcc_path = os.path.dirname(os.path.dirname(os.path.realpath(__file__)))
        self.filepath = os.path.join(etcc_path,'frontend/index.html')
    def on_get(self, req, resp):
        if (os.path.isfile(self.filepath)):
            resp.status = falcon.HTTP_200
            resp.stream = open(self.filepath, 'rb')
            resp.stream_len = os.path.getsize(self.filepath)
            resp.content_type = 'text/html'
        else:
            resp.status = falcon.HTTP_404
