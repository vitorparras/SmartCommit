import { Component } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { TranslateService } from '@ngx-translate/core';

@Component({
  selector: 'app-api-key-modal',
  templateUrl: './api-key-modal.component.html',
  styleUrls: ['./api-key-modal.component.scss']
})
export class ApiKeyModalComponent {
  apiKey: string = '';
  showKey: boolean = false;

  constructor(
    public dialogRef: MatDialogRef<ApiKeyModalComponent>,
    public translate: TranslateService
  ) {
    // Remove default dialog padding
    this.dialogRef.addPanelClass('api-key-dialog');
  }

  toggleVisibility() {
    this.showKey = !this.showKey;
  }

  onSubmit() {
    if (this.apiKey.trim()) {
      this.dialogRef.close(this.apiKey);
    }
  }

  onCancel() {
    this.dialogRef.close();
  }
}

