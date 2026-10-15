"use strict";

// Minimal leveled logger replacing simple-node-logger, which is unmaintained
// and crashes on Node >= 23 (it relies on util.isDate and other legacy util
// predicates that were removed). The transformer only needs a
// console-compatible object whose output is gated by a minimum level.

// Capture the real console up front: cli.js swaps global.console for the
// logger, so looking it up lazily would recurse.
const realConsole = {
  debug: console.debug.bind(console),
  info: console.info.bind(console),
  warn: console.warn.bind(console),
  error: console.error.bind(console),
};

// Ordered from most to least verbose. "all" is the lowest threshold; trace and
// fatal are distinct levels, not aliases of debug and error.
const LEVELS = ["all", "trace", "debug", "info", "warn", "error", "fatal"];
const DEFAULT_LEVEL = "warn";
const noop = () => {};

// Each method name is also its level, so a method is enabled when its level is
// at or above the configured minimum.
const SINKS = {
  trace: realConsole.debug,
  debug: realConsole.debug,
  info: realConsole.info,
  warn: realConsole.warn,
  error: realConsole.error,
  fatal: realConsole.error,
};

function createLogger(level) {
  const logger = {
    setLevel(newLevel) {
      let min = LEVELS.indexOf(newLevel);
      if (min === -1) {
        min = LEVELS.indexOf(DEFAULT_LEVEL);
      }
      Object.entries(SINKS).forEach(([name, sink]) => {
        logger[name] = LEVELS.indexOf(name) >= min ? sink : noop;
      });
      // console-compatible alias
      logger.log = logger.info;
    },
  };
  logger.setLevel(level || DEFAULT_LEVEL);
  return logger;
}

module.exports = { createLogger };
