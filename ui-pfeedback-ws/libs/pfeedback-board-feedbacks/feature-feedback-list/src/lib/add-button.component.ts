import { RouterModule } from '@angular/router';
import {
  Component,
  EventEmitter,
  Input,
  NgModule,
  Output,
} from '@angular/core';

@Component({
  selector: 'app-add-button',
  template: `
    <a
      [routerLink]="[link]"
      class="rounded-lg px-6 py-3 bg-fuchsia-700 text-white font-bold text-sm flex items-center space-x-1"
    >
      <ng-content></ng-content>
    </a>
  `,
})
export class AddButtonComponent {
  @Output()
  clickAction: EventEmitter<void> = new EventEmitter<void>();

  @Input()
  link!: string;
}

@NgModule({
  imports: [RouterModule],
  declarations: [AddButtonComponent],
  exports: [AddButtonComponent],
})
export class AddButtonModule {}
