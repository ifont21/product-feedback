import { CommonModule } from '@angular/common';
import { Component, Input, NgModule } from '@angular/core';

@Component({
  selector: 'app-ui-badge',
  template: `
    <div class="bg-zinc-100 rounded-lg text-sm text-blue-700 font-semibold py-2 px-4 inline-flex">
      <!-- <span class="text-sm tracking-tight font-bold text-indigo-900"> -->
        <ng-content></ng-content>
      <!-- </span> -->
    </div>
  `,
})
export class UIBadgeComponent {
  @Input()
  votes = 0;
}

@NgModule({
  imports: [CommonModule],
  declarations: [UIBadgeComponent],
  exports: [UIBadgeComponent],
})
export class UIBadgeModule {}
