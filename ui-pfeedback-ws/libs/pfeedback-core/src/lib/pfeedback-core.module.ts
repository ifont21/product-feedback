import { ModuleWithProviders, NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SpecConfig as BoardSpecConfig } from '@pfeedback/board-data-spec';
import { SpecConfig as FeedbackSpecConfig } from '@pfeedback/feedback-data-spec';

@NgModule({
  imports: [CommonModule],
})
export class PFeedbackCoreModule {
  static forRoot({
    serverUrl,
  }: {
    serverUrl: string;
  }): ModuleWithProviders<PFeedbackCoreModule> {
    return {
      ngModule: PFeedbackCoreModule,
      providers: [
        {
          provide: BoardSpecConfig,
          useValue: {
            basePath: `${serverUrl}/api`,
          },
        },
        {
          provide: FeedbackSpecConfig,
          useValue: {
            basePath: `${serverUrl}/api`,
          },
        },
      ],
    };
  }
}
