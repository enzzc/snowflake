# Lock-free Snowflake ID generator

This is a lock-free implementation of an algorithm that generates
Snowflake identifiers thread-safely and without a mutex.

This version returns IDs structured as follows:

 * High 48 bits: Unix timestamp (in milliseconds)
 * Low 16 bits: Sequence number (0 to 65535)

This is similar to Mastodon's version, but
different from Twitter's original version.

This implementation seems to work, but it should not be considered
production-ready. Actually, this is the code that illustrates
my note about lock-free programming: <https://lu.sagebl.eu/notes/implementing-lock-free-snowflake-id-generator/>
