
# Approach:

As opposed to trying to parse with a default, then fallback to a list of options,
receive a list of configurations directly. That way, there's no unnecessary code
being run, except for configuration error. Which is highly performant, and
somewhat documentable, since the configuration will tell directly what should happen in the parsing phase.

# TO DO
* Don't vendor the proto definitions, and instead refer directly to the `opentelemetry-proto` directory.
* Create a custom CLI to send synthetic data instead of using a script
* Receive the configuration via a file and maybe via the cli?
* Support more types than just 'json' and 'regex', maybe gzip and other binary formats? Or custom formats like logfmt (key value) etc
* Create a pipeline for transformations, so that one could plug regex + json parser
* Deal with different encodings
* ParserPipeline should be moved to Parser package and then implement the Parse interface
* When flushing, print in the ingestion order
* Don't convert `<nil>` to unknown, instead handle it appropriately

## Brainstorm
Create a cli to generate synthetic data.

* Should we use map for everything, or for the fields we know specify them.

* Hardcode the config file, since it's not in scope to implement a fully functional cli/config loader.

* Should we log unknowns? In one hand, they look like wrong values, on the other hand they could be totally expected absences. Maybe put under verbose log level?

* Should we use opentelemetry collector code? Not because it's not the point of
the exercise, although may be worth looking into it for a production application (at least for a MVP).

# Ideas:
* Instead of running the main code for configured options object code under a goroutine, move to parts that are highly concurrent.


# Approach
1. Construct a simple script to generate gRPC data

# References

The `.proto` file -> https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/collector/logs/v1/logs_service.proto
