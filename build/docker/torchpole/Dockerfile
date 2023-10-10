# Copyright 2022 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/combizent/torchpole.

FROM BASE_IMAGE
LABEL maintainer="<dairongpeng@foxmail.com>"

WORKDIR /opt/torchpole

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone && \
      mkdir -p /var/log/torchpole

COPY torchpole /opt/torchpole/bin/

ENTRYPOINT ["/opt/torchpole/bin/torchpole"]
CMD ["-c", "/etc/torchpole/torchpole.yaml"]
