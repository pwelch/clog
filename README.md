# Clog

Clog is a CLI client written in Go for the [Captain's Log](https://github.com/pwelch/captains_log) Rails application.

## Configuration

Clog looks for a configuration file `/etc/clog/clog_config.json` by default.

To specify a different file use `-c /path/to/config`.

Example clog configuration file

```bash
# clog_config.json

{
 "server": "http://example.com/api",
 "api_key": "API_KEY"
}
```

## Usage

```bash
  clog -m "Restarted apache in staging"
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
