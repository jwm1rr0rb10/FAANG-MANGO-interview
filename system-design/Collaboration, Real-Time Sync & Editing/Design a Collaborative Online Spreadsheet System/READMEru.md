# Дизайн Коллаборативной Онлайн-Таблицы

## Overview

![Overview1](https://github.com/ogamor69wm1rr0rb/senior_question_interview/blob/main/images/backend/system-design/Collaboration%2C%20Real-Time%20Sync%20%26%20Editing/Design%20a%20Collaborative%20Online%20Spreadsheet%20System/overview1.png)
![Overview1](https://github.com/ogamor69wm1rr0rb/senior_question_interview/blob/main/images/backend/system-design/Collaboration%2C%20Real-Time%20Sync%20%26%20Editing/Design%20a%20Collaborative%20Online%20Spreadsheet%20System/overview2.png)
![Overview1](https://github.com/ogamor69wm1rr0rb/senior_question_interview/blob/main/images/backend/system-design/Collaboration%2C%20Real-Time%20Sync%20%26%20Editing/Design%20a%20Collaborative%20Online%20Spreadsheet%20System/overview3.jpeg)
![Overview1](https://github.com/ogamor69wm1rr0rb/senior_question_interview/blob/main/images/backend/system-design/Collaboration%2C%20Real-Time%20Sync%20%26%20Editing/Design%20a%20Collaborative%20Online%20Spreadsheet%20System/overview4.jpeg)

Real-time collaborative editing systems позволяют нескольким пользователям одновременно редактировать один документ и видеть изменения почти мгновенно.

Типичные примеры:

## Google Docs

Google Docs — это бесплатный облачный текстовый редактор, разработанный компанией Google. Он позволяет отдельным пользователям и командам создавать, редактировать документы и совместно работать над ними в режиме реального времени с любого устройства, имеющего доступ в интернет. Интегрированный в Google Workspace, он является краеугольным камнем онлайн-продуктивности и сотрудничества для миллионов пользователей по всему миру.- [Google-Docs](https://www.britannica.com/topic/Google-Docs?utm_source=chatgpt.com)￼

### Ключевые факты

- Год запуска: 2006 (создано на основе приобретенного сервиса Writely)
- Платформа: Web, Android, iOS, ChromeOS
- Основной пакет: Google Workspace
- Искусственный интеллект-помощник: Gemini в Docs (ранее Duet AI)
- Цены: Бесплатно для частных лиц; премиум-версия через тарифные планы Workspace

---

### Истоки и развитие

- Google Docs возник на основе Writely, раннего браузерного текстового редактора, созданного компанией Upstartle в 2005 году. Google приобрела его в 2006 году, объединив с недавно запущенным пакетом Google Apps (позже переименованным в Google Workspace). Со временем Docs превратился из простого онлайн-редактора в сложную облачную платформу, конкурирующую с настольными текстовыми редакторами, такими как Microsoft Word.[techtarget](https://www.techtarget.com/whatis/definition/Google-Docs?utm_source=chatgpt.com)
---

### Сотрудничество и функциональность

- Docs allows multiple users to edit the same file simultaneously, with real-time cursor tracking and colour-coded changes. A built-in version history logs every revision, enabling users to review or restore past edits. Collaboration tools include comments, suggestions, task assignments, and in-document chat. Offline mode permits continued editing without internet connectivity, syncing automatically once reconnected.  ￼

---

### Features and integrations

Docs supports common file types including .docx, .pdf, and .txt, allowing seamless import and export. Templates, pageless documents, smart chips, and voice typing improve productivity. Integration with other Google apps—Drive, Gmail, Sheets, and Meet—lets users embed charts, reply to comments via email, or launch video calls directly within documents.  ￼

---

### AI-powered writing and smart tools

Modern versions of Docs include AI features under Gemini, Google’s generative-AI suite. Users can prompt Gemini to draft, refine, or summarise text, generate structured documents, and access contextual “help me write” suggestions. Additional assistive tools such as Smart Compose, autocorrect, and built-in summaries enhance writing fluency and efficiency.  ￼

---

### Security and access

Documents are encrypted in transit and at rest, managed through user-defined sharing permissions. Workspace administrators gain enterprise-grade controls such as client-side encryption, access logs, and compliance certifications. Users need only a Google Account to access the free version, while business and education tiers offer extended features and support.  ￼

---

## Notion

---

## Figma

---

## Dropbox Paper

Ключевая задача:

поддерживать консистентное состояние документа при одновременных изменениях от множества пользователей.

Основные требования:
•	latency < 100 ms
•	высокая консистентность документа
•	устойчивость к сетевым задержкам
•	масштабирование на миллионы одновременных пользователей
