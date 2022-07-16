import { CommonModule } from '@angular/common';
import {
  Component,
  EventEmitter,
  Input,
  NgModule,
  Output,
} from '@angular/core';
import { DropdownOption } from './types';

@Component({
  selector: 'app-ui-dropdown',
  template: ` <div class="relative">
    <div
      tabindex="0"
      (click)="toggleOpen()"
      class="h-12 w-full p-5 flex justify-between items-center rounded bg-slate-100 cursor-pointer"
    >
      <span>{{ selected ? selected : 'Pick One' }}</span>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="16"
        height="16"
        fill="currentColor"
        class="bi bi-chevron-up"
        viewBox="0 0 16 16"
      >
        <path
          fill-rule="evenodd"
          d="M7.646 4.646a.5.5 0 0 1 .708 0l6 6a.5.5 0 0 1-.708.708L8 5.707l-5.646 5.647a.5.5 0 0 1-.708-.708l6-6z"
        />
      </svg>
    </div>
    <ul
      *ngIf="open"
      class="w-full absolute mt-4 rounded-lg top-full bg-white shadow-2xl"
    >
      <li
        *ngFor="let option of options"
        class="px-6 py-3 border-b border-indigo-900/10 cursor-pointer"
        (click)="selectItem(option)"
      >
        {{ option.label }}
      </li>
    </ul>
  </div>`,
})
export class UIDropdownComponent {
  @Input()
  options: DropdownOption[] = [];

  @Output()
  selectedItem: EventEmitter<string | number> = new EventEmitter<
    string | number
  >();

  selected: string | number = '';
  open = false;

  toggleOpen(): void {
    this.open = !this.open;
  }

  selectItem(value: DropdownOption): void {
    this.selected = value.label;
    this.selectedItem.emit(value.key);
    this.toggleOpen();
  }
}

@NgModule({
  imports: [CommonModule],
  declarations: [UIDropdownComponent],
  exports: [UIDropdownComponent],
})
export class UIDropdownModule {}
