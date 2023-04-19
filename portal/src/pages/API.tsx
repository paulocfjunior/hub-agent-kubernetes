/*
Copyright (C) 2022-2023 Traefik Labs
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.
You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

import 'components/styles/css-root.css'
import 'components/styles/swagger-ui.css'
import 'components/styles/swagger-fixes.css'
import 'components/styles/css-root-override.css'
import React, { useMemo } from 'react'
import { Box } from '@traefiklabs/faency'
import { useParams } from 'react-router-dom'
import { Helmet } from 'react-helmet-async'
import SwaggerUI from 'swagger-ui-react'

import { AugmentedLayoutPlugin } from 'components/layouts/AugmentedLayout'
import { getInjectedValues } from 'utils/getInjectedValues'

// const requestInterceptor = (req) => {
//   const token = localStorage.getItem('token')
//   return {
//     ...req,
//     headers: {
//       Authorization: `Bearer ${token}`,
//     },
//   }
// }

const API = () => {
  const { portalName } = getInjectedValues()
  const { apiName, collectionName } = useParams()

  const specUrl = useMemo(() => {
    if (collectionName) {
      return `/api/${portalName}/collections/${collectionName}/apis/${apiName}`
    }

    return `/api/${portalName}/apis/${apiName}`
  }, [collectionName, portalName, apiName])

  return (
    <Box>
      <Helmet>
        <title>{apiName || 'API Portal'}</title>
      </Helmet>
      <Box>
        <SwaggerUI layout="AugmentedLayout" plugins={[AugmentedLayoutPlugin]} url={specUrl} />
      </Box>
    </Box>
  )
}

export default API
