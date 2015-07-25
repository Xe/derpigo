# db-massdelete
--
Command db-massdelete handles mass deletion of derpibooru images.

Please use this sparingly.

Usage:

    Usage of ./db-massdelete:
      -keyFile="/home/xena/.local/share/within/db.cadance.key": file with the derpibooru key to use
      -reason="": reason to use when deleting images

Then give it the image ID's you want to delete.

    ./db-massdelete -reason "OP is a duck" 123 325 1561 136324
