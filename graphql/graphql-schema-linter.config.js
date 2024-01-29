module.exports = {
  "rules": [
    "defined-types-are-used",
    "deprecations-have-a-reason",
    "descriptions-are-capitalized",
    "enum-values-all-caps",
    "fields-are-camel-cased",
    "input-object-values-are-camel-cased",
    "interface-fields-sorted-alphabetically",
    "relay-connection-arguments-spec",
    "relay-connection-types-spec",
    "types-are-capitalized"

    // 可能な限り名前で役割を分かるようにする方針のため、descriptionは強制しない。
    // "arguments-have-descriptions"
    // "enum-values-have-descriptions"
    // "fields-have-descriptions"
    // "input-object-values-have-descriptions"
    // "types-have-descriptions"

    // ページネーションが必要になった時に有効化する。
    // "relay-page-info-spec"

    // 意味のある順番にしたいため無効化。例えばstartDateとendDateは連続させたいなど。
    // "enum-values-sorted-alphabetically"
    // "input-object-fields-sorted-alphabetically"
    // "type-fields-sorted-alphabetically"
  ],
  "schemaPaths": [
    "*.graphqls"
  ]
}
