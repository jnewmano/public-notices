package locationbing

const sampleResponse = `{
  "authenticationResultCode": "ValidCredentials",
  "brandLogoUri": "http://dev.virtualearth.net/Branding/logo_powered_by.png",
  "copyright": "Copyright © 2017 Microsoft and its suppliers. All rights reserved. This API cannot be accessed and the content and any results may not be used, reproduced or transmitted in any manner without express written permission from Microsoft Corporation.",
  "resourceSets": [
    {
      "estimatedTotal": 1,
      "resources": [
        {
          "__type": "Location:http://schemas.microsoft.com/search/local/ws/rest/v1",
          "bbox": [
            40.4081922352685,
            -111.883500128112,
            40.4159176704098,
            -111.869972464691
          ],
          "name": "2051 N 1450 W, Lehi, UT 84043",
          "point": {
            "type": "Point",
            "coordinates": [
              40.4120549528392,
              -111.876736296402
            ]
          },
          "address": {
            "addressLine": "2051 N 1450 W",
            "adminDistrict": "UT",
            "adminDistrict2": "Utah",
            "countryRegion": "United States",
            "formattedAddress": "2051 N 1450 W, Lehi, UT 84043",
            "locality": "Lehi",
            "postalCode": "84043"
          },
          "confidence": "Medium",
          "entityType": "Address",
          "geocodePoints": [
            {
              "type": "Point",
              "coordinates": [
                40.4120549528392,
                -111.876736296402
              ],
              "calculationMethod": "InterpolationOffset",
              "usageTypes": [
                "Display"
              ]
            },
            {
              "type": "Point",
              "coordinates": [
                40.4120717640033,
                -111.876681522092
              ],
              "calculationMethod": "Interpolation",
              "usageTypes": [
                "Route"
              ]
            }
          ],
          "matchCodes": [
            "Good"
          ]
        }
      ]
    }
  ],
  "statusCode": 200,
  "statusDescription": "OK",
  "traceId": "abc|abc|7.7.0.0|"
}`
