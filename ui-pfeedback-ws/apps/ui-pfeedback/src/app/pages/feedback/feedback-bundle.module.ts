import { RouterModule } from '@angular/router';
import { Component, NgModule } from '@angular/core';

@Component({
  selector: 'app-feedback-bundle',
  template: ` <div>Feedback Page!</div> `,
})
export class FeedbackBundleComponent {}

const routes = [
  {
    path: '',
    component: FeedbackBundleComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  declarations: [FeedbackBundleComponent]
})
export class FeedbackBundleModule {}
