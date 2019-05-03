#!/bin/bash

function cleanup_xml_roundtrip {
    for i in $(seq 0 1 $2); do
        CURRENT_FILE="${1}/${i}"
        CURRENT_FILE_ROUNDTRIP="${CURRENT_FILE}.roundtrip.xml"
        CURRENT_FILE_JSON="${CURRENT_FILE}.json"

        xmllint --c14n "${CURRENT_FILE_ROUNDTRIP}" > "${CURRENT_FILE}.tmp";
        xmllint --pretty 1 "${CURRENT_FILE}.tmp" > "${CURRENT_FILE_ROUNDTRIP}";
        rm "${CURRENT_FILE}.tmp";

        jq "." "${CURRENT_FILE_JSON}" > "${CURRENT_FILE}.tmp"
        mv "${CURRENT_FILE}.tmp" "${CURRENT_FILE_JSON}"
    done
}

function cleanup_json_roundtrip {
    for i in $(seq 0 1 $2); do
        CURRENT_FILE="${1}/${i}"
        CURRENT_FILE_ROUNDTRIP="${CURRENT_FILE}.roundtrip.json"
        CURRENT_FILE_JSON="${CURRENT_FILE}.json"

        jq "." "${CURRENT_FILE_ROUNDTRIP}" > "${CURRENT_FILE}.tmp"
        mv "${CURRENT_FILE}.tmp" "${CURRENT_FILE_ROUNDTRIP}"

        jq "." "${CURRENT_FILE_JSON}" > "${CURRENT_FILE}.tmp"
        mv "${CURRENT_FILE}.tmp" "${CURRENT_FILE_JSON}"
    done
}

function cleanup_json {
    for i in $(seq 0 1 $2); do
        CURRENT_FILE="${1}/${i}"
        CURRENT_FILE_JSON="${CURRENT_FILE}.json"

        jq "." "${CURRENT_FILE_JSON}" > "${CURRENT_FILE}.tmp"
        mv "${CURRENT_FILE}.tmp" "${CURRENT_FILE_JSON}"
    done
}

CURRENT_DIR=$(dirname $(realpath "$0"))

cleanup_xml_roundtrip "${CURRENT_DIR}/iris_station" 2
cleanup_xml_roundtrip "${CURRENT_DIR}/iris_realtime" 4
cleanup_xml_roundtrip "${CURRENT_DIR}/iris_timetable" 4
cleanup_xml_roundtrip "${CURRENT_DIR}/iris_wingdef" 2
cleanup_json_roundtrip "${CURRENT_DIR}/apps_wagenreihung" 226
cleanup_json "${CURRENT_DIR}/hafas_messages" 24
