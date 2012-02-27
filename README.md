# Golaroid - Image filtering web-service in Go 

Golaroid is a simple web-service which provides picture filtering.

## Installation

Use the `go install` tool:

    $ go install github.com/nu7hatch/golaroid

## Usage

Start golaroid service on given address and pointed to specified images
root location:

    $ golaroid -addr=':8090' -image-root=./

Now go to `http://127.0.0.1:8090/bubble.jpg?filter=sepia` to see effects.

## Filters

So far project is on very early states and supports only the following
filters:

* desaturate
* sepia

## Copyright

Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
