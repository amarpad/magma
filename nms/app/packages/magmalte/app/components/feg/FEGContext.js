/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @flow strict-local
 * @format
 */

import * as React from 'react';
import FEGNetworkContext from '../context/FEGNetworkContext';
import LoadingFiller from '@fbcnms/ui/components/LoadingFiller';
import MagmaV1API from '@fbcnms/magma-api/client/WebClient';
import useMagmaAPI from '@fbcnms/ui/magma/useMagmaAPI';

import type {feg_network, network_id, network_type} from '@fbcnms/magma-api';

import {UpdateNetworkState as UpdateFegNetworkState} from '../../state/feg/NetworkState';
import {useCallback, useState} from 'react';
import {useEnqueueSnackbar} from '@fbcnms/ui/hooks/useSnackbar';

type Props = {
  networkId: network_id,
  networkType: network_type,
  children: React.Node,
};

/**
 * Fetches and returns information about the federation network inside
 * a context provider.
 * @param {object} props: contains the network id and its type
 */
export function FEGNetworkContextProvider(props: Props) {
  const {networkId} = props;
  const [fegNetwork, setFegNetwork] = useState<$Shape<feg_network>>({});
  const enqueueSnackbar = useEnqueueSnackbar();
  const {error, isLoading} = useMagmaAPI(
    MagmaV1API.getFegByNetworkId,
    {networkId: networkId},
    useCallback(response => setFegNetwork(response), []),
  );

  if (error) {
    enqueueSnackbar?.('failed fetching network information', {
      variant: 'error',
    });
  }

  if (isLoading) {
    return <LoadingFiller />;
  }

  return (
    <FEGNetworkContext.Provider
      value={{
        state: fegNetwork,
        updateNetworks: props => {
          let refreshState = true;
          if (networkId !== props.networkId) {
            refreshState = false;
          }
          return UpdateFegNetworkState({
            networkId,
            setFegNetwork,
            refreshState,
            ...props,
          });
        },
      }}>
      {props.children}
    </FEGNetworkContext.Provider>
  );
}
/**
 * A context provider for federation networks. It is used in sharing
 * information like the network information or the gateways information.
 * @param {object} props contains the network id and its type
 */
export function FEGContextProvider(props: Props) {
  const {networkId, networkType} = props;

  return (
    <FEGNetworkContextProvider {...{networkId, networkType}}>
      {props.children}
    </FEGNetworkContextProvider>
  );
}
