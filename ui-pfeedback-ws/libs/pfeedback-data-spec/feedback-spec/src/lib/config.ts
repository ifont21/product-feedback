import { InjectionToken } from '@angular/core';

export interface Configuration {
  basePath: string;
}

export const SpecConfig = new InjectionToken<Configuration>(
  'Base path configuration'
);
