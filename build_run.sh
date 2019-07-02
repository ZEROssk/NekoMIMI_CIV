#!/bin/bash
docker build ./ -t img_v/zero:latest
docker run --name img_v_zero -it --rm -p 5300:5300 img_v/zero:latest
