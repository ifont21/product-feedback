import { PFeedbackHeaderContainerModule } from '@pfeedback/shell/header';
import { CommonModule } from '@angular/common';
import { Component, NgModule } from '@angular/core';

@Component({
  selector: 'app-pfeedback-shell',
  template: `
    <div class="h-screen flex flex-col">
      <!-- Header Navbar -->
      <app-pfeedback-header-container></app-pfeedback-header-container>
      <!-- Router outlet Body content-->
      <div class="grow">
        <div class="max-w-[80%] my-20 mx-auto">
          <!-- <app-auth-button></app-auth-button>
            <app-projects></app-projects> -->
          <ng-content></ng-content>
        </div>
      </div>
    </div>
  `,
})
export class PFeedbackShellComponent {}

@NgModule({
  imports: [CommonModule, PFeedbackHeaderContainerModule],
  declarations: [PFeedbackShellComponent],
  exports: [PFeedbackShellComponent],
})
export class PFeedbackShellModule {}
