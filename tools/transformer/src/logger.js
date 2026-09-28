"use strict";

// Minimal leveled logger replacing simple-node-logger, which is unmaintained
// and crashes on Node >= 23 (it relies on util.isDate and other legacy util
// predicates that were removed). The transformer only needs a
// console-compatible object whose debug/info/warn/error output is gated by a
// minimum level.

// Capture the real console up front: cli.js swaps global.console for the
// logger, so looking it up lazily would recurse.
const realConsole = {
  debug: console.debug.bind(console),
  info: console.info.bind(console),
  warn: console.warn.bind(console),
  error: console.error.bind(console),
};

const LEVELS = ["debug", "info", "warn", "error"];
const DEFAULT_LEVEL = "warn";
const noop = () => {};

function createLogger(level) {
  const logger = {
    setLevel(newLevel) {
      let min = LEVELS.indexOf(newLevel);
      if (min === -1) {
        min = LEVELS.indexOf(DEFAULT_LEVEL);
      }
      LEVELS.forEach((name, index) => {
        logger[name] = index >= min ? realConsole[name] : noop;
      });
      // console-compatible aliases
      logger.log = logger.info;
      logger.trace = logger.debug;
      logger.fatal = logger.error;
    },
  };
  logger.setLevel(level || DEFAULT_LEVEL);
  return logger;
}

module.exports = { createLogger };
