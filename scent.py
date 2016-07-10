import time
from sniffer.api import *
from subprocess import call
from os.path import basename
import termstyle

pass_fg_color = termstyle.green
pass_bg_color = termstyle.bg_default
fail_fg_color = termstyle.red
fail_bg_color = termstyle.bg_default

watch_paths = ["."]


def throttle(validator):
    def inner(filename):
        if inner.last < time.time() - 3:
            inner.last = time.time()
            return validator(filename)
        else:
            return False

    inner.last = time.time()
    return inner


@select_runnable('go_tests')
@file_validator
@throttle
def go_files(filename):
    return filename.endswith('.go') and not basename(filename).startswith('.')


@runnable
def go_tests(*args):
    return call("go test github.com/aychedee/seer", shell=True) == 0
