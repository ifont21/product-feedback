import { RouterModule } from '@angular/router';
import { Component, NgModule } from '@angular/core';

@Component({
  selector: 'app-roadmap-bundle',
  template: ` <div>Road-map View Page!</div> `,
})
export class RoadMapViewBundleComponent {}

const routes = [
  {
    path: '',
    component: RoadMapViewBundleComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  declarations: [RoadMapViewBundleComponent],
})
export class RoadMapViewBundleModule {}
