"""
Logging module for bost
"""
import logging
import sys

logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.INFO, stream=sys.stdout)
