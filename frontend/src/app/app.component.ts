import { Component, OnInit } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { MatDialog } from '@angular/material/dialog';
import { LanguageModalComponent } from './shared/components/language-modal/language-modal.component';

interface Language {
  code: string;
  name: string;
  nativeName: string;
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  languages: Language[] = [
    { code: 'us', name: 'English', nativeName: 'English' },
    { code: 'es', name: 'Spanish', nativeName: 'Español' },
    { code: 'fr', name: 'French', nativeName: 'Français' },
    { code: 'de', name: 'German', nativeName: 'Deutsch' },
    { code: 'br', name: 'Portuguese', nativeName: 'Português' }
  ];

  constructor(
    public translate: TranslateService,
    private dialog: MatDialog
  ) {
    translate.setDefaultLang('us');
    translate.addLangs(['us', 'es', 'fr', 'de', 'br']);
  }

  ngOnInit() {
    const savedLang = localStorage.getItem('selectedLanguage');
    if (savedLang) {
      this.translate.use(savedLang);
    } else {
      this.openLanguageModal();
    }
  }

  openLanguageModal() {
    const dialogRef = this.dialog.open(LanguageModalComponent, {
      width: '400px',
      maxWidth: '90vw',
      panelClass: 'language-modal',
      disableClose: true
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.translate.use(result);
        localStorage.setItem('selectedLanguage', result);
      }
    });
  }

  getCurrentLanguageName(): string {
    const currentLang = this.translate.currentLang || 'us';
    const language = this.languages.find(lang => lang.code === currentLang);
    return language ? language.nativeName : 'English';
  }

  handleImageError(event: Event): void {
    const img = event.target as HTMLImageElement;
    img.style.display = 'none';
  }
}

