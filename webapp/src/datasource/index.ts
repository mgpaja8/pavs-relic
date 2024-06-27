import { Datasource, DatasourceConfig } from './pavsRelic/Datasource';

export * from './pavsRelic/Interface';

const datasourceConfig: DatasourceConfig = {
  baseURL: process.env.REACT_APP_SERVER_URL ?? ''
};
export const pavsRelicDatasource = new Datasource(datasourceConfig);
