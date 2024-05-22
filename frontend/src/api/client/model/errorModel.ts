/**
 * Media Recylarr
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 0.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { ErrorDetail } from './errorDetail';

export class ErrorModel {
    /**
    * A URL to the JSON Schema for this object.
    */
    '$schema'?: string;
    /**
    * A human-readable explanation specific to this occurrence of the problem.
    */
    'detail'?: string;
    /**
    * Optional list of individual error details
    */
    'errors'?: Array<ErrorDetail>;
    /**
    * A URI reference that identifies the specific occurrence of the problem.
    */
    'instance'?: string;
    /**
    * HTTP status code
    */
    'status'?: number;
    /**
    * A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
    */
    'title'?: string;
    /**
    * A URI reference to human-readable documentation for the error.
    */
    'type'?: string = 'about:blank';

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "$schema",
            "baseName": "$schema",
            "type": "string"
        },
        {
            "name": "detail",
            "baseName": "detail",
            "type": "string"
        },
        {
            "name": "errors",
            "baseName": "errors",
            "type": "Array<ErrorDetail>"
        },
        {
            "name": "instance",
            "baseName": "instance",
            "type": "string"
        },
        {
            "name": "status",
            "baseName": "status",
            "type": "number"
        },
        {
            "name": "title",
            "baseName": "title",
            "type": "string"
        },
        {
            "name": "type",
            "baseName": "type",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return ErrorModel.attributeTypeMap;
    }
}

